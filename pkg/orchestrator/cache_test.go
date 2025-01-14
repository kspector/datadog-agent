// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package orchestrator

import (
	"testing"

	model "github.com/DataDog/agent-payload/v5/process"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/types"
)

func TestKubeCacheHits(t *testing.T) {
	uid := "123-456-789"
	resourceVersion := "123"
	newResourceVersion := "321"
	// we do not have the resource, therefore we do not skip this resource.
	skip := SkipKubernetesResource(types.UID(uid), resourceVersion, model.K8SResource_POD)
	assert.False(t, skip)

	// we have the resource, therefore we skip this resource.
	skip = SkipKubernetesResource(types.UID(uid), resourceVersion, model.K8SResource_POD)
	assert.True(t, skip)

	// we have the resource but the version has changed. therefore we do not skip this resource.
	skip = SkipKubernetesResource(types.UID(uid), newResourceVersion, model.K8SResource_POD)
	assert.False(t, skip)
}
