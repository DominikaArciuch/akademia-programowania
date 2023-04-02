package academy

import "math"

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) > 0 {
		suma := 0.0
		for _, value := range grades {
			suma += float64(value)
		}
		return int(math.RoundToEven(suma / float64(len(grades))))
	}
	return 0

}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) > 0 {
		suma := 0
		for _, value := range attendance {
			if value {
				suma++
			}
		}
		return math.Round(float64(suma)/float64(len(attendance))*1000) / 1000
	}
	return 0.00
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	grade := int(math.Round((float64(s.Project) + float64(AverageGrade(s.Grades))) / 2.0))
	if AttendancePercentage(s.Attendace) < 0.6 || s.Project == 1 || AverageGrade(s.Grades) == 1 {
		return 1
	} else if AttendancePercentage(s.Attendace) < 0.8 && grade > 1 {
		return grade - 1
	} else {
		return grade
	}
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	var value = map[string]uint8{}
	for _, student := range students {
		value[student.Name] = uint8(FinalGrade(student))
	}
	return value
}
