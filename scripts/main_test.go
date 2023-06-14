package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_IN_PATH = "../nicknames.txt"
const TEST_OUT_PATH = "../nicknames.json"

func TestGetNicknames(t *testing.T) {
	nicknames, err := getNicknames(TEST_IN_PATH)
	assert.Nil(t, err)
	assert.True(t, len(*nicknames) > 0)
}

func TestConvertToNicknameJPENList(t *testing.T) {
	nicknameJPENList := convertToNicknameJPENList(&[]string{"りんご", "みかん"})
	arr := *nicknameJPENList
	assert.Equal(t, 2, len(arr))
	assert.Equal(t, "りんご", arr[0][0])
	assert.Equal(t, "ringo", arr[0][1])
	assert.Equal(t, "みかん", arr[1][0])
	assert.Equal(t, "mikan", arr[1][1])
}

func TestSaveJson(t *testing.T) {
	err := saveJson(TEST_OUT_PATH, &[][]string{})
	assert.Nil(t, err)
}
