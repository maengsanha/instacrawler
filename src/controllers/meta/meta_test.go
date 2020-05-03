// Package meta implements Instagram meta-search engine logics.
package meta

import "fmt"

func Example_meta_Search() {
	secondLayerQueries := []string{"cafe", "카페", "coffee", "커피", "일상"}
	thirdLayerQueries := []string{"starbucks", "스타벅스", "daily", "korea", "coffee"}
	secondLayerCache := []string{"", "", "", "", ""}
	thirdLayerCache := []string{"", "", "", "", ""}
	resp := Search(secondLayerQueries, thirdLayerQueries, secondLayerCache, thirdLayerCache)
	fmt.Println(len(resp.SecondLayer))
	fmt.Println(len(resp.ThirdLayer))
	fmt.Println(resp.SecondLayerCache)
	fmt.Println(resp.ThirdLayerCache)
	// Output:
	// 606
	// 319
	// [QVFEVFlfdFdVcVZkZnk1aVU4QlFGSnI4eV8taWRraGZGS1FxRGR3OXI1TGNSaFdtUF84TnNmMzV0eXB1UnZ0R1lTczdjXzBQcFZUVlIzNUs3NjBIZ1lMQQ== QVFDeWliSHZKUXVxVkg2amNNT0xSYUZqWmhYS3ZhTGw3R2lnLXBkQU9QUjRjbTBfMEs4S21aS3VicFpSMjVRT1FUb0dmQUhMTUw1a2wydlpCblNwY3gxcg== QVFCZjRxYTdGSXNNV2ZHbXhhUG5VZVVaRURWWEVxNmtvb2FpWnItZVRUeTBFUGRLNElqV0wwMG5CVWlISlBkYnZuWWRyQkt6VzllbGl4ejZVZWM4R09RRw== QVFCYkNiczZhUkpCdnI4NVF2YzFDeVJ2YzNKUHJ4R2NpZFpjckJ0ZzdPZmM0b20tZ1ljRHV3Rk56U1ZKekRVZVAtZjRybnduaWdYSTNkMWtsczRneVE3Zw== QVFDZG4wMU5oVHc4N0VGdm9yQWhYM0JBLXM0VEMwTWdNVl9Na3JleGxCc181N0VZTUlIR0VJbkxNWHZyN1dfREJNQ2FKN3hGTzlUU1ZRWURwanh2SVdKeg==]
	// [QVFDZHJia09TNEhmWTlid0w0dlpqemQ5MWY3UUpIQmtqQXQwWTJFOEdwSXBSWE5DT25QeHY2RXF6LWF4LUEwOHpJdDgwamRaNnJGTUZnQWhJSnRvYk8tYg== QVFDdkhzYzBaQVB0VE1QY041ZVdIVTJBQ2NMdkVKRWVqMWo3bDZDby0wVHRlTFM1bE5NTS15dTZHR2twOGZ2VERkNGhZQjh0S0RjNTV0QlhaMlViNTF6Uw== QVFEQzhMRzE3RmxNZlRERkdzOHQwV0gzZHVqZDJjY2VBVFhzbnJrZ3B6a3dpVXpmNGQ4T1ppODVabndtZ2tSRkxtN2hKV3cxZ0FaLVhwZEdMOFdUanlmMQ== QVFBT2RRZF9hZUZjMmJnRXMtbmR6X3Y3QTFOcXMwYjNxRGN1djRBUDJFeWR3elNCOFpPNE5pMlVScjBoc3E0dGtsZjE1Z2lpZTdhZDh4QmxQdHlUUmJXdQ== QVFETG9XUVNvZ1dWOFRMT1RZQ2diNjJGczIycnE0YnJlZjc2MGFmUFV0ZDB6NjBDSXVfeE13VW1tVHJlTHRCeXZhWjlUazJsUzhHcWxxbjNVRUNBR1ZIVg==]
}
