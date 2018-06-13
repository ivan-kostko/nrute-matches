package application

import (
	"context"
	"strconv"

	"github.com/ivan-kostko/nrute-matches/domain"
)

func MatchMovementsToBundleContractConditions(ctx context.Context, movements []Movement, conds []domain.ContractCondition) []Match {

	mainLogger := new(log).WithFields(map[string]interface{}{"logger": "MatchMovementsToBundleContractConditions"})
	mainLogger.Info("MatchMovementsToBundleContractConditions invoked")

	mainLogger.Debug("Getting all combinations")

	combinations := getMatchingCombinations(mainLogger, movements, conds)

	mainLogger.Debug("Selecting the best from combinations")

	theBest := selectBestMatchCombination(mainLogger, combinations)

	if len(theBest) == 0 {
		mainLogger.Info("No (best)matche(s) found. The best is just unmatched movements")
		theBest = []Match{Match{Movements: movements}}
	}

	return theBest
}

func selectBestMatchCombination(logger Log, combinations [][]Match) []Match {

	if len(combinations) == 0 {
		logger.Info("No combinations provided for selecting the best one. Returning")
		return nil
	}

	logger.Debug("Selecting the best combination from:", combinations)

	winners := struct {
		BestScore    int
		Combinations [][]Match
	}{}

	for combinationNo, combination := range combinations {

		combinationLogger := logger.WithFields(map[string]interface{}{"combination_no": combinationNo})

		combinationLogger.Info("Selecting best combination")
		combinationLogger.Debug("Combination score set to zero")
		combinationScore := 0

		for _, match := range combination {

			matchLogIdentifier := "Unmatched movements"
			if match.ContractCondition != nil {
				matchLogIdentifier = match.ContractCondition.Id
			}
			matchLogger := combinationLogger.WithFields(map[string]interface{}{"match_identifier": matchLogIdentifier})

			matchLogger.Debug("Combination score seed for " + strconv.Itoa(match.Score))
			combinationScore += match.Score
			matchLogger.Debug("Current combination score is " + strconv.Itoa(combinationScore))
		}

		if combinationScore == winners.BestScore {
			combinationLogger.Debug("Current combination has same score as some in before. Adding to potential winner(s)")
			winners.Combinations = append(winners.Combinations, combination)

		}

		if combinationScore > winners.BestScore {
			combinationLogger.Debug("Current combination is better than any in before. Selecting as potential winner")
			winners.Combinations = [][]Match{combination}
			winners.BestScore = combinationScore

		}

	}

	if len(winners.Combinations) > 1 {
		logger.WithFields(map[string]interface{}{"winners_best_score": winners.BestScore}).Warn("More than one combination has the best score")
		return nil
	}

	// There should be one-and-only-one combination, cause casewith 0 combinations was excluded in the beginning of the func

	logger.WithFields(map[string]interface{}{"winners_best_score": winners.BestScore}).Info("The winner successfully selected")
	logger.Debug("The winner is: ", winners.Combinations[0])
	return winners.Combinations[0]

}

func getMatchingCombinations(logger Log, movements []Movement, conds []domain.ContractCondition) [][]Match {

	logger.Info("getMatchingCombinations invoked with the following params:\r\n", movements, conds)

	const (
		VehicleTypeDirectMatchScore              = 3
		VehicleTypeFallbackMatchScore            = 0
		WorkflowFactorDirectMatchScore           = 2
		WorkflowFactorFallbackMatchScore         = 0
		MovementActivityOptionDirectMatchScore   = 1
		MovementActivityOptionFallbackMatchScore = 0
	)

	// Represents set of match combinations as the result returned by this function.
	// It will be appended on every success iteration.
	resultMatchCombinations := [][]Match{}

	for condNo, cond := range conds {

		condLogger := logger.WithFields(map[string]interface{}{"contract_condition_id": cond.Id, "contract_condition_name": cond.Name})

		// Skipping non bundles
		if !(len(cond.MovementActivities) > 1) {
			condLogger.Info("ContractCondition is skipped because number of movement activities (" + strconv.Itoa(len(cond.MovementActivities)) + ") is less than 2, so it is not a Bundle at all.")
			continue
		}

		// Represents set of movements which have not been matched yet.
		// Due to rearrangements of `unmatchedMovementLeftovers` it is better to copy original `movements`.
		unmatchedMovementLeftovers := append([]Movement{}, movements...)

		// Skip CC if it has more MA than movements at all.
		// It wont match anyway...
		if len(cond.MovementActivities) > len(unmatchedMovementLeftovers) {
			condLogger.Debug("ContractCondition is skipped because number of movement activities (" + strconv.Itoa(len(cond.MovementActivities)) + ") is more than number of matching movements(" + strconv.Itoa(len(unmatchedMovementLeftovers)) + "), so wont match at all.")
			continue
		}

		condLogger.Info("Starting to find matches for contract condition")

		// The flag indicates whether current contract condition has movement activity w/o corresponding movement.
		ccHasUnmatchedMA := false

		condLogger.Debug("Movements to exercise: ", unmatchedMovementLeftovers)

		// Allocate a new copy of cond and reference to it while &cond - will point to variable cond which is getting new content on each iteration.
		// So, &cond at the end on this function will point to cond variable, which will have content of last iterated contract condition.
		conditionCopy := cond

		// In case of successfull matching it will be appended to resultMatchCombinations at the very end of iteration.
		// Meanwhile, on each MA 2 MVMT match it will be updated.
		currentCcMatch := Match{ContractCondition: &conditionCopy}

		// Loop over CCs MovementActivities
		for _, ccma := range cond.MovementActivities {

			ccmaLogger := condLogger.WithFields(map[string]interface{}{"movement_activity_type": ccma.Type, "movement_activity_option": ccma.Option})
			ccmaLogger.Debug("Starting to match movement to current activity")

			// Lets be objective - it is not matched yet.
			maHasMatched := false

			// Loop over `movements` in order to find matches
			for mvmtNo, mvmt := range unmatchedMovementLeftovers {

				mvmtLogger := ccmaLogger.WithFields(map[string]interface{}{"movement_id": mvmt.Id})
				mvmtLogger.Info("Matching movement to CC MA")

				// Extract contractor identifier from movement.
				// It is needed later to check if movement fits cc.
				mvmtContractorId := ""
				if mvmt.User.Contractor != nil {
					mvmtContractorId = *(mvmt.User.Contractor)
				}

				doesnotMatch := false
				switch {
				case cond.ContractorIdentifier != mvmtContractorId:
					mvmtLogger.Debug("Movement ContractorId (" + mvmtContractorId + "does not match CC ContractorIdentifier (" + cond.ContractorIdentifier + ")")
					doesnotMatch = true
					fallthrough
				case cond.BranchIdentifier != mvmt.Branch.Id:
					mvmtLogger.Debug("Movement Branch.Id (" + mvmt.Branch.Id + ") does not match CC BranchIdentifier (" + cond.BranchIdentifier + ")")
					doesnotMatch = true
					fallthrough
				case cond.WorkflowType != mvmt.Workflow.Type:
					mvmtLogger.Debug("Movement Workflow.Type (" + mvmt.Workflow.Type + "does not match CC WorkflowType (" + cond.WorkflowType + ")")
					doesnotMatch = true
					fallthrough
				case ccma.Type != mvmt.Type:
					mvmtLogger.Debug("Movement Type (" + mvmt.Type + "does not match CC MA Type (" + ccma.Type + ")")
					doesnotMatch = true
					// Add more checks like for CC validity date, etc...
				}
				// Skip if doesn't match.
				if doesnotMatch {
					mvmtLogger.Debug("Movement does not match by main properties")
					continue
				}

				// Check for VehicleType match
				switch {
				case cond.VehicleType == mvmt.Vehicle.Type:
					mvmtLogger.Debug("Movement Vehicle.Type directly matches to contract condition VehicleType")
					currentCcMatch.Score += VehicleTypeDirectMatchScore
				case cond.VehicleType == domain.Undefined_VehicleType:
					mvmtLogger.Debug("Movement Vehicle.Type matches to fallback")
					currentCcMatch.Score += VehicleTypeFallbackMatchScore
				default:
					// This contract condition wont match, cause VehicleType does not match neither movement nor fallback
					mvmtLogger.Debug("Movement VehicleType does not match neither ContractCondition nor fallback. Movement is skipped")
					continue
				}

				// Check for WorkflowFactor match
				switch {
				case cond.WorkflowFactor == mvmt.Workflow.Factor:
					mvmtLogger.Debug("Movement Workflow.Factor directly matches to contract condition WorkflowFactor")
					currentCcMatch.Score += WorkflowFactorDirectMatchScore
				case cond.WorkflowFactor == domain.Undefined_WorkflowFactor:
					mvmtLogger.Debug("Movement  Workflow.Factor matches to fallback")
					currentCcMatch.Score += WorkflowFactorFallbackMatchScore
				default:
					// This contract condition wont match, cause VehicleType does not match neither movement nor fallback
					mvmtLogger.Debug("Movement WorkflowFactor does not match neither ContractCondition nor fallback. Movement is skipped")
					continue
				}

				// Check for Option match
				switch {
				case ccma.Option == mvmt.Option:
					mvmtLogger.Debug("Movement Option directly matches to contract condition movement activity option")
					currentCcMatch.Score += MovementActivityOptionDirectMatchScore
				case ccma.Option == domain.Undefined_MovementOption:
					mvmtLogger.Debug("Movement  Option matches to fallback")
					currentCcMatch.Score += MovementActivityOptionFallbackMatchScore
				default:
					// This contract condition wont match, cause VehicleType does not match neither movement nor fallback
					mvmtLogger.Debug("Movement Option does not match neither ContractCondition nor fallback. Movement is skipped")
					continue
				}

				mvmtLogger.Info("Movement matched to contract condition movement activity. Appending it to current match and removing it from unmatched.")

				// Remove current move from unmatched
				unmatchedMovementLeftovers = append(unmatchedMovementLeftovers[:mvmtNo], unmatchedMovementLeftovers[mvmtNo+1:]...)

				mvmtLogger.Debug("Unmatched leftovers: ", unmatchedMovementLeftovers)

				// Append current match to movements
				currentCcMatch.Movements = append(currentCcMatch.Movements, mvmt)

				// Indicate that MA has matched.
				maHasMatched = true

				break
			}

			if !maHasMatched {
				// Means no movement matches MA - deal with it!
				ccmaLogger.Info("Noone movement matches movement activity")
				ccHasUnmatchedMA = true
				break
			}

		}

		if ccHasUnmatchedMA {
			condLogger.Info("Skipping contract condition while there are unmatched movement activities")
			continue
		}

		condLogger.Debug("Current ContractCondition match: ", currentCcMatch)

		// All previously checked contract conditions will appear in resultMatchCombinations if matched.
		// So, only further/leftover conditions should be checked for matching to unmatched movements
		conditionLeftovers := conds[condNo+1:]

		condLogger.Info("Contract Condition matched to some movements. Calling to match the following leftovers: ", unmatchedMovementLeftovers, conditionLeftovers)

		leftoverCombinations := [][]Match{}

		if len(unmatchedMovementLeftovers) > 0 && len(conditionLeftovers) > 0 {
			condLogger.Info("Tere are movements and CCs left. Calling to match leftovers")
			leftoverCombinations = getMatchingCombinations(logger, unmatchedMovementLeftovers, conditionLeftovers)
		}

		condLogger.Debug("Leftover combinations are as the following: ", leftoverCombinations)

		if len(leftoverCombinations) > 0 {
			condLogger.Debug("There are matched leftovers. Adding them with current match to result")
			for _, leftoverCombination := range leftoverCombinations {
				// Appending resultMatchCombinations with leftoverCombination in conjunction with current match
				resultMatchCombinations = append(resultMatchCombinations, append(leftoverCombination, currentCcMatch))
			}

			// Due to fact that any match is better than unmatched movements, ther is no reason to add current match in conjunction to unmatched movements.
			// Also, it could bring problems in case of same score with some match and zero score. So, just moving to next CC.

			continue
		}

		condLogger.Debug("There are no matched leftovers.")

		// While there were no matched leftovers, for consistency it should return current match with unmatched movements.
		currentCombination := []Match{currentCcMatch}
		if len(unmatchedMovementLeftovers) > 0 {

			condLogger.Debug("Adding unmatched movements to current match combination")

			currentCombination = append(currentCombination, Match{Movements: unmatchedMovementLeftovers})
		}

		resultMatchCombinations = append(resultMatchCombinations, currentCombination)

	}

	return resultMatchCombinations
}
