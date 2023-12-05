package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getWholeNumber(t *testing.T) {
	type cases struct {
		s              string
		startIndex     int
		expectedNumber int64
		expectedIndex  int
	}

	testCases := []cases{
		{
			s:              "467..114..",
			startIndex:     0,
			expectedNumber: 467,
			expectedIndex:  2,
		},
		{
			s:              "467..114..",
			startIndex:     1,
			expectedNumber: 467,
			expectedIndex:  2,
		},
		{
			s:              "467..114..",
			startIndex:     6,
			expectedNumber: 114,
			expectedIndex:  7,
		},
		{
			s:              "467..114..",
			startIndex:     7,
			expectedNumber: 114,
			expectedIndex:  7,
		},
		{
			s:              "254...@............$.......@..#17..............284.........757......................921..#479..933.414....308......@154............*.....222",
			startIndex:     137,
			expectedNumber: 222,
			expectedIndex:  139,
		},
		
	}
	for _, tc := range testCases {
		actualNumber, actualIndex := getWholeNumber(tc.s, tc.startIndex)
		require.Equal(t, tc.expectedNumber, actualNumber)
		require.Equal(t, tc.expectedIndex, actualIndex)
	}
}
