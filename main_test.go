package main

import (
    "testing"
)

func TestMaximum(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected int
    }{
        {
            name:     "empty slice should return 0",
            input:    []int{},
            expected: 0,
        },
        {
            name:     "single element",
            input:    []int{42},
            expected: 42,
        },
        {
            name:     "all positive numbers",
            input:    []int{1, 2, 3, 4, 5},
            expected: 5,
        },
        {
            name:     "all negative numbers",
            input:    []int{-5, -4, -3, -2, -1},
            expected: -1,
        },
        {
            name:     "mixed positive and negative",
            input:    []int{-10, 0, 10, -20, 20},
            expected: 20,
        },
        {
            name:     "with duplicates",
            input:    []int{3, 3, 3, 3, 3},
            expected: 3,
        },
        {
            name:     "large numbers",
            input:    []int{1<<31 - 1, -1 << 31, 0},
            expected: 1<<31 - 1,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := maximum(tt.input)
            if result != tt.expected {
                t.Errorf("maximum(%v) = %d, want %d", tt.input, result, tt.expected)
            }
        })
    }
}