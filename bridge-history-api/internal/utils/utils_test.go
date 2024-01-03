package utils

import (
	"testing"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestKeccak2(t *testing.T) {
	a := common.HexToHash("0xe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0")
	b := common.HexToHash("0x222ff5e0b5877792c2bc1670e2ccd0c2c97cd7bb1672a57d598db05092d3d72c")
	c := Keccak2(a, b)
	assert.NotEmpty(t, c)
	assert.NotEqual(t, a, c)
	assert.NotEqual(t, b, c)
	assert.Equal(t, "0xc0ffbd7f501bd3d49721b0724b2bff657cb2378f15d5a9b97cd7ea5bf630d512", c.Hex())
}

func TestGetBatchRangeFromCalldata(t *testing.T) {
	// single chunk
	start, finish, err := GetBatchRangeFromCalldata(common.Hex2Bytes("1325aca000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000005900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003d0100000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000100000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	assert.Equal(t, start, uint64(1))
	assert.Equal(t, finish, uint64(1))

	// multiple chunk
	start, finish, err = GetBatchRangeFromCalldata(common.Hex2Bytes("1325aca000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000003e0000000000000000000000000000000000000000000000000000000000000007900000000000000000100000000000000010000000000000001038433daac85a0b03cd443ed50bc85e832c883061651ae2182b2984751e0b340119b828c2a2798d2c957228ebeaff7e10bb099ae0d4e224f3eeb779ff61cba610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000004c01000000000000000a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000010000000001000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001b403000000000000000b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000300000000000000000b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00050000000000000014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012c01000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000100000000010000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa800000000000000000000000000000000000000000000000000000000000000aa"))
	assert.NoError(t, err)
	assert.Equal(t, start, uint64(10))
	assert.Equal(t, finish, uint64(20))

	// genesis batch
	start, finish, err = GetBatchRangeFromCalldata(common.Hex2Bytes("3fdeecb200000000000000000000000000000000000000000000000000000000000000402dcb5308098d24a37fc1487a229fcedb09fa4343ede39cbad365bc925535bb09000000000000000000000000000000000000000000000000000000000000005900000000000000000000000000000000000000000000000000c252bc9780c4d83cf11f14b8cd03c92c4d18ce07710ba836d31d12da216c8330000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	assert.NoError(t, err)
	assert.Equal(t, start, uint64(0))
	assert.Equal(t, finish, uint64(0))
}