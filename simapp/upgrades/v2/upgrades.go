
package v2 // upgrades.go

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// var DenomMetadata = banktypes.Metadata{...}
func CreateUpgradeHandler(
    mm *module.Manager,
    configurator module.Configurator,
) upgradetypes.UpgradeHandler {
    return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
    
       // UpdateDenomMetadata(ctx, bk)
        
        return mm.RunMigrations(ctx, configurator, vm)
    }
}

/* func UpdateDenomMetadata(ctx sdk.Context, bk bankkeeper.Keeper) {
    bk.SetDenomMetaData(ctx, DenomMetadata)
} */
