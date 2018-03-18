package application_test

import (
	"context"
	"testing"
	"time"

	"github.com/ivan-kostko/nrute-matches/application"
	"github.com/ivan-kostko/nrute-matches/domain"

	"github.com/stretchr/testify/assert"
)

func TestMatchMovementsToBundleContractConditions(t *testing.T) {
	testCases := []struct {
		Alias           string
		MovementsIn     []application.Movement
		ConditionsIn    []domain.ContractCondition
		ExpectedMatches []application.Match
	}{
		/*
			Test cases 4 beer
		*/
		{
			Alias: `First Andy case`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{

				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "",
								Type:   "parking",
							},
						},
					},
					Score: 8,
				},
			},
		},
		{
			Alias: `Second Andy case`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "vip",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "TheFirst",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
				domain.ContractCondition{
					Id:                   "JustSecond",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "vip",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{

				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "vip",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "JustSecond",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "standard",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "vip",
								Type:   "parking",
							},
						},
					},
					Score: 12,
				},
			},
		},
		{
			Alias: `Third Andy case`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132455",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132456",
					Type:     "refueling",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "98765"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "refueling",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{

				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132455",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "standard",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "",
								Type:   "parking",
							},
						},
					},
					Score: 12,
				},
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "refueling",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "98765"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
				},
			},
		},

		/*
			Test cases for code coverage
		*/

		{
			Alias: `Movements do not match by main properties to 1 CC with 2 MA`,
			MovementsIn: []application.Movement{
				// wrong Contractor id
				application.Movement{
					Id:       "132450",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "1231465", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "wrong_987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				// wrong branch
				application.Movement{
					Id:       "132451",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "wrong_6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				// wrong Workflow type
				application.Movement{
					Id:       "132452",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314655", Type: "wrong_turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "someOpt",
							Type:   "checkin",
						},
						{
							Option: "someOpt",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{

				application.Match{
					Movements: []application.Movement{
						// wrong Contractor id
						application.Movement{
							Id:       "132450",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "1231465", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "wrong_987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// wrong branch
						application.Movement{
							Id:       "132451",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "wrong_6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// wrong Workflow type
						application.Movement{
							Id:       "132452",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314655", Type: "wrong_turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: nil,
					Score:             0,
				},
			},
		},
		{
			Alias: `Movements do not match by sub properties to 1 CC with 2 MA`,
			MovementsIn: []application.Movement{

				// wrong Vehicle type
				application.Movement{
					Id:       "132453",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "wrong_car", Id: "TheVehicleId"},
				},
				// wrong WF factor
				application.Movement{
					Id:       "132454",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "wrong_standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				// wrong option
				application.Movement{
					Id:       "132455",
					Type:     "checkin",
					Option:   "wrong_",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "someOpt",
							Type:   "checkin",
						},
						{
							Option: "someOpt",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{

				application.Match{
					Movements: []application.Movement{

						// wrong Vehicle type
						application.Movement{
							Id:       "132453",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "wrong_car", Id: "TheVehicleId"},
						},
						// wrong WF factor
						application.Movement{
							Id:       "132454",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "wrong_standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// wrong option
						application.Movement{
							Id:       "132455",
							Type:     "checkin",
							Option:   "wrong_",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: nil,
					Score:             0,
				},
			},
		},
		{
			Alias: `2 Movements match to 1 CC with 2 MA`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "standard",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "",
								Type:   "parking",
							},
						},
					},
					Score: 12,
				},
			},
		},
		{
			Alias: `5 Movements match to 1 CC with 2 MA`,
			MovementsIn: []application.Movement{

				// tpye="exterior_cleaning" is not in CCMAs at all
				application.Movement{
					Id:       "132453",
					Type:     "exterior_cleaning",
					Option:   "automative",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				// tpye="interior_cleaning" is not in CCMAs at all
				application.Movement{
					Id:       "132454",
					Type:     "interior_cleaning",
					Option:   "wet cleaning",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				// tpye="interior_cleaning" is not in CCMAs at all
				application.Movement{
					Id:       "132455",
					Type:     "interior_cleaning",
					Option:   "dry-cleaning",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "standard",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "",
								Type:   "parking",
							},
						},
					},
					Score: 12,
				},
				application.Match{
					Movements: []application.Movement{

						// tpye="exterior_cleaning" is not in CCMAs at all
						application.Movement{
							Id:       "132453",
							Type:     "exterior_cleaning",
							Option:   "automative",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// tpye="interior_cleaning" is not in CCMAs at all
						application.Movement{
							Id:       "132454",
							Type:     "interior_cleaning",
							Option:   "wet cleaning",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// tpye="interior_cleaning" is not in CCMAs at all
						application.Movement{
							Id:       "132455",
							Type:     "interior_cleaning",
							Option:   "dry-cleaning",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: nil,
					Score:             0,
				},
			},
		},
		{
			Alias: `7 Movements match to many CCs`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132451",
					Type:     "refuel",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				application.Movement{
					Id:       "132452",
					Type:     "refill_watertank",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				application.Movement{
					Id:       "132453",
					Type:     "exterior_cleaning",
					Option:   "bikini_girls",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132454",
					Type:     "interior_cleaning",
					Option:   "wet cleaning",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				application.Movement{
					Id:       "132455",
					Type:     "interior_cleaning",
					Option:   "dry-cleaning",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "CC-1",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
				domain.ContractCondition{
					Id:                   "CC-2",
					Name:                 "SuperCleaning",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "",
					VehicleType:          "",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "exterior_cleaning",
						},
						{
							Option: "dry-cleaning",
							Type:   "interior_cleaning",
						},
						{
							Option: "wet cleaning",
							Type:   "interior_cleaning",
						},
					},
				},
				// Should be skipped because there are less movements
				domain.ContractCondition{
					Id:                   "CC-3",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "",
					VehicleType:          "",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
						{
							Option: "",
							Type:   "",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132453",
							Type:     "exterior_cleaning",
							Option:   "bikini_girls",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},

						application.Movement{
							Id:       "132455",
							Type:     "interior_cleaning",
							Option:   "dry-cleaning",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						application.Movement{
							Id:       "132454",
							Type:     "interior_cleaning",
							Option:   "wet cleaning",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "CC-2",
						Name:                 "SuperCleaning",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "",
						VehicleType:          "",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "exterior_cleaning",
							},
							{
								Option: "dry-cleaning",
								Type:   "interior_cleaning",
							},
							{
								Option: "wet cleaning",
								Type:   "interior_cleaning",
							},
						},
					},

					Score: 2,
				},
				application.Match{
					Movements: []application.Movement{

						// Type not in any CC
						application.Movement{
							Id:       "132451",
							Type:     "refuel",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						// Type not in any CC
						application.Movement{
							Id:       "132452",
							Type:     "refill_watertank",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: nil,
					Score:             0,
				},
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
					ContractCondition: &domain.ContractCondition{
						Id:                   "CC-1",
						Name:                 "Turnaround",
						WorkflowType:         "turnaround",
						WorkflowFactor:       "standard",
						VehicleType:          "car",
						BranchIdentifier:     "6",
						ContractorIdentifier: "987654",
						MovementActivities: []domain.MovementActivity{
							{
								Option: "",
								Type:   "checkin",
							},
							{
								Option: "",
								Type:   "parking",
							},
						},
					},
					Score: 12,
				},
			},
		},
	}

	for _, tCase := range testCases {

		testFn := func(t *testing.T) {

			ctx := context.Background()

			actualMatches := application.MatchMovementsToBundleContractConditions(ctx, tCase.MovementsIn, tCase.ConditionsIn)

			if !(assert.Subset(t, tCase.ExpectedMatches, actualMatches) ||
				assert.Subset(t, actualMatches, tCase.ExpectedMatches)) {
				t.Log("Returned Combination: \r\n")

				for mNo, m := range actualMatches {
					t.Logf(" MatchNo: %d \r\n%#v\r\n", mNo, m)

				}

			}
		}

		t.Run(tCase.Alias, testFn)
	}

}

func TestMatchMovementsToBundleContractConditions_OnSameScore(t *testing.T) {

	testCases := []struct {
		Alias           string
		MovementsIn     []application.Movement
		ConditionsIn    []domain.ContractCondition
		ExpectedMatches []application.Match
	}{
		{
			Alias: `2 Movements match to 2 CC with 2 MA`,
			MovementsIn: []application.Movement{
				application.Movement{
					Id:       "132456",
					Type:     "checkin",
					Option:   "option1",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},

				application.Movement{
					Id:       "132457",
					Type:     "parking",
					Option:   "option2",
					Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
					Branch:   application.Branch{Id: "6"},
					Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
					User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
					Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
				},
			},
			ConditionsIn: []domain.ContractCondition{
				domain.ContractCondition{
					Id:                   "WF+Opts",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "standard",
					VehicleType:          "",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "option1",
							Type:   "checkin",
						},
						{
							Option: "option2",
							Type:   "parking",
						},
					},
				},
				domain.ContractCondition{
					Id:                   "VT",
					Name:                 "Turnaround",
					WorkflowType:         "turnaround",
					WorkflowFactor:       "",
					VehicleType:          "car",
					BranchIdentifier:     "6",
					ContractorIdentifier: "987654",
					MovementActivities: []domain.MovementActivity{
						{
							Option: "",
							Type:   "checkin",
						},
						{
							Option: "",
							Type:   "parking",
						},
					},
				},
			},
			ExpectedMatches: []application.Match{
				application.Match{
					Movements: []application.Movement{
						application.Movement{
							Id:       "132456",
							Type:     "checkin",
							Option:   "option1",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
						application.Movement{
							Id:       "132457",
							Type:     "parking",
							Option:   "option2",
							Date:     time.Date(2018, 01, 31, 16, 59, 59, 999999990, time.UTC),
							Branch:   application.Branch{Id: "6"},
							Workflow: application.Workflow{Id: "12314654", Type: "turnaround", Factor: "standard"},
							User:     application.User{Contractor: func() *string { s := "987654"; return &s }(), Id: "TheUserId"},
							Vehicle:  application.Vehicle{Type: "car", Id: "TheVehicleId"},
						},
					},
				},
			},
		},
	}

	for _, tCase := range testCases {

		testFn := func(t *testing.T) {

			ctx := context.Background()

			actualMatches := application.MatchMovementsToBundleContractConditions(ctx, tCase.MovementsIn, tCase.ConditionsIn)

			if !(assert.Subset(t, tCase.ExpectedMatches, actualMatches) ||
				assert.Subset(t, actualMatches, tCase.ExpectedMatches)) {
				t.Log("Returned Combination: \r\n")

				for mNo, m := range actualMatches {
					t.Logf(" MatchNo: %d \r\n%#v\r\n", mNo, m)

				}

			}
		}

		t.Run(tCase.Alias, testFn)
	}

}
