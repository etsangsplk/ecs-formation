package bluegreen

import (
	"github.com/openfresh/ecs-formation/service"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Show plan to execute bluegreen deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		bgsrv, err := service.NewBlueGreenService(projectDir, bluegreenName, parameters)
		if err != nil {
			return err
		}

		csrv, err := bgsrv.CreateClusterService()
		if err != nil {
			return err
		}

		if _, err := createBlueGreenPlans(bgsrv, csrv); err != nil {
			return err
		}

		return nil
	},
}
