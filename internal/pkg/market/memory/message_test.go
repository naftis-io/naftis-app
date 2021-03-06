package memory

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gitlab.com/naftis/app/naftis/pkg/protocol/market"
	"testing"
	"time"
)

func TestMessage_ListenWorkloadSpecification(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.WorkloadSpecification{}

	ch1 := m.ListenWorkloadSpecification(ctx, 3)
	assert.Empty(t, ch1)

	ch2 := m.ListenWorkloadSpecification(ctx, 3)
	assert.Empty(t, ch2)

	txId, err = m.EmitWorkloadSpecification(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 1)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 1)

	ch3 := m.ListenWorkloadSpecification(ctx, 3)
	assert.Empty(t, ch3)

	txId, err = m.EmitWorkloadSpecification(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 2)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 2)
	assert.NotEmpty(t, ch3)
	assert.Len(t, ch3, 1)
}

func TestMessage_ListenContractAccept(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.ContractAccept{}

	ch1 := m.ListenContractAccept(ctx, 3)
	assert.Empty(t, ch1)

	ch2 := m.ListenContractAccept(ctx, 3)
	assert.Empty(t, ch2)

	txId, err = m.EmitContractAccept(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 1)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 1)

	ch3 := m.ListenContractAccept(ctx, 3)
	assert.Empty(t, ch3)

	txId, err = m.EmitContractAccept(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 2)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 2)
	assert.NotEmpty(t, ch3)
	assert.Len(t, ch3, 1)
}

func TestMessage_ListenContractProposal(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.ContractProposal{}

	ch1 := m.ListenContractProposal(ctx, 3)
	assert.Empty(t, ch1)

	ch2 := m.ListenContractProposal(ctx, 3)
	assert.Empty(t, ch2)

	txId, err = m.EmitContractProposal(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 1)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 1)

	ch3 := m.ListenContractProposal(ctx, 3)
	assert.Empty(t, ch3)

	txId, err = m.EmitContractProposal(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	// Wait for message delivery, dirty but working
	time.Sleep(time.Millisecond)

	assert.NotEmpty(t, ch1)
	assert.Len(t, ch1, 2)
	assert.NotEmpty(t, ch2)
	assert.Len(t, ch2, 2)
	assert.NotEmpty(t, ch3)
	assert.Len(t, ch3, 1)
}

func TestMessage_EmitContractProposal(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.ContractProposal{}

	txId, err = m.EmitContractProposal(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitContractProposal(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitContractProposal(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)
}

func TestMessage_EmitContractAccept(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.ContractAccept{}

	txId, err = m.EmitContractAccept(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitContractAccept(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitContractAccept(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)
}

func TestMessage_EmitWorkloadSpecification(t *testing.T) {
	var err error
	var txId string

	ctx := context.TODO()
	m := NewMessage()
	m.Start(ctx)

	msg := market.WorkloadSpecification{}

	txId, err = m.EmitWorkloadSpecification(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitWorkloadSpecification(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)

	txId, err = m.EmitWorkloadSpecification(ctx, msg)
	assert.NotEmpty(t, txId)
	assert.NoError(t, err)
}
