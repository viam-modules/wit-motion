// package main is a module with raspberry pi board component.
package main

import (
	"context"

	"github.com/viam-modules/wit-motion/imuwit"

	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("wit-motion"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	module, err := module.NewModuleFromArgs(ctx)
	if err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, movementsensor.API, imuwit.Model); err != nil {
		return err
	}

	if err = module.AddModelFromRegistry(ctx, movementsensor.API, imuwit.Model905); err != nil {
		return err
	}

	err = module.Start(ctx)
	defer module.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
