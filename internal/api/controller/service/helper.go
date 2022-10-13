// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package service

import (
	"context"

	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/types"
)

func findServiceFromUID(ctx context.Context,
	serviceStore store.ServiceStore, serviceUID string) (*types.Service, error) {
	return serviceStore.FindUID(ctx, serviceUID)
}
