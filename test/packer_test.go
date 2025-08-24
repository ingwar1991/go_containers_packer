package test

import (
    "testing"
    "runtime"
    "time"
    "containers_packer/internal/packer"
)


type TestCase struct {
    Containers []int
    Goods int
    Result map[int]int 
}

var cases = []TestCase{
    {
        []int{2, 7, 9},
        18,
        map[int]int{
            9: 2,
        },
    },
    {
        []int{7, 9},
        5,
        map[int]int{
            7: 1,
        },
    },
    {
        []int{2, 7},
        14,
        map[int]int{
            7: 2,
        },
    },
    {
        []int{2, 7, 9},
        20,
        map[int]int{
            2: 1,
            9: 2,
        },
    },
    {
        []int{4, 9},
        15,
        map[int]int{
            4: 4,
        },
    },
    {
        []int{6, 9, 20},
        43,
        map[int]int{
            6: 1,
            9: 2,
            20: 1,
        },
    },
    {
        []int{5, 7, 13},
        100,
        map[int]int{
            5: 3,
            7: 1,
            13: 6,
        },
    },
    {
        []int{2, 50},
        200,
        map[int]int{
            50: 4,
        },
    },
    {
        []int{7, 11},
        20,
        map[int]int{
            7: 3,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        1,
        map[int]int{
            250: 1,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        250,
        map[int]int{
            250: 1,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        251,
        map[int]int{
            500: 1,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        501,
        map[int]int{
            250: 1,
            500: 1,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        4300,
        map[int]int{
            500: 1,
            2000: 2,
        },
    },
    {
        []int{250, 500, 1000, 2000, 5000},
        12001,
        map[int]int{
            250: 1,
            2000: 1,
            5000: 2,
        },
    },
    {
        []int{23, 31, 53},
        500000,
        map[int]int{
            23: 2,
            31: 7,
            53: 9429,
        },
    },
}

func mapsAreEqual(map1, map2 map[int]int) bool {
    for key1, val1 := range map1 {
        if val2, ok := map2[key1]; !ok || val1 != val2 {
            return false
        }
    }

    return true
}

func Test(t *testing.T) {
    var mStart, mEnd runtime.MemStats

    for _, testCase := range cases {
        start := time.Now()
        runtime.ReadMemStats(&mStart)

        actual, err := packer.Pack(testCase.Containers, testCase.Goods)
        if err != nil {
            t.Errorf("Error during exec: %v", err)
        }
        
        if !mapsAreEqual(actual, testCase.Result) {
            containersCtnActual, emptySpaceActual := 0, 0
            for size, cnt := range actual {
                emptySpaceActual += size * cnt
                containersCtnActual += cnt
            }
            emptySpaceActual -= testCase.Goods

            containersCtn, emptySpace := 0, 0
            for size, cnt := range testCase.Result {
                emptySpace += size * cnt
                containersCtn += cnt
            }
            emptySpace -= testCase.Goods

            if emptySpace != emptySpaceActual || containersCtn != containersCtnActual {
                t.Errorf("Expected result: %v, but got %v for %v, %v", testCase.Result, actual, testCase.Containers, testCase.Goods)
                continue
            }

            t.Logf("Found new valid result: %v, vs expected %v for %v, %v", actual, testCase.Result, testCase.Containers, testCase.Goods)
        }

        runtime.ReadMemStats(&mEnd)
        allocated := mEnd.TotalAlloc - mStart.TotalAlloc

        duration := time.Since(start)

        t.Logf("%v\n\tExecution time: %v\n\tMemory allocated: %d bytes", testCase, duration, allocated)

    }
}
