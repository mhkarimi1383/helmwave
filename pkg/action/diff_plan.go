package action

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/urfave/cli/v2"
)

type DiffPlans struct {
	plandir1, plandir2 string
	diff               *Diff
}

func (d *DiffPlans) Run() error {
	if d.plandir1 == d.plandir2 {
		log.Warn(plan.ErrPlansAreTheSame)
	}

	plan1 := plan.New(d.plandir1)
	if err := plan1.Import(); err != nil {
		return err
	}
	if ok := plan1.IsManifestExist(); !ok {
		return os.ErrNotExist
	}

	plan2 := plan.New(d.plandir2)
	if err := plan2.Import(); err != nil {
		return err
	}
	if ok := plan2.IsManifestExist(); !ok {
		return os.ErrNotExist
	}

	plan1.DiffPlan(plan2, d.diff.ShowSecret, d.diff.Wide)

	return nil
}

func (d *DiffPlans) Cmd() *cli.Command {
	return &cli.Command{
		Name:   "plan",
		Usage:  "plan1  🆚  plan2",
		Flags:  d.flags(),
		Action: toCtx(d.Run),
	}
}

func (d *DiffPlans) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "plandir1",
			Value:       ".helmwave/",
			Usage:       "Path to plandir1",
			EnvVars:     []string{"HELMWAVE_PLANDIR_1", "HELMWAVE_PLANDIR"},
			Destination: &d.plandir1,
		},
		&cli.StringFlag{
			Name:        "plandir2",
			Value:       ".helmwave/",
			Usage:       "Path to plandir2",
			EnvVars:     []string{"HELMWAVE_PLANDIR_2"},
			Destination: &d.plandir2,
		},
	}
}
