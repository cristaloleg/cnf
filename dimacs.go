package cnf

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Problem struct {
	Variables int
	Clauses   int
	Formula   Formula
}

func ParseDIMAC(r io.Reader) (*Problem, error) {
	result := &Problem{
		Variables: -1,
	}

	var current []Lit

	scanner := bufio.NewScanner(r)
	read := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if result.Variables == -1 {
			switch line[0] {
			case 'c':
				continue

			case 'p':
				fields := strings.Fields(line)
				if len(fields) != 4 {
					return nil, fmt.Errorf("problem line should have 4 fields whitespace separated: %q", line)
				}

				if fields[1] != "cnf" {
					return nil, fmt.Errorf("problem type must be 'cnf', got: %q", fields[1])
				}

				vars, err := strconv.Atoi(fields[2])
				if err != nil {
					return nil, fmt.Errorf("error converting variable count %q: %w", fields[2], err)
				}

				clauses, err := strconv.Atoi(fields[3])
				if err != nil {
					return nil, fmt.Errorf("error converting clauses count %q: %w", fields[3], err)
				}

				result.Variables = vars
				result.Clauses = clauses
				continue

			default:
				return nil, fmt.Errorf("invalid start of line character: %q", line[0])
			}
		}

		fields := strings.Fields(line)

		var end bool

		for _, raw := range fields {
			val, err := strconv.Atoi(raw)
			if err != nil {
				return nil, fmt.Errorf("invalid literal: %q", raw)
			}

			if val == 0 {
				end = true
				break
			}
			current = append(current, NewLit(val))
		}
		if !end {
			continue
		}

		slices.Sort(current)

		result.Formula = append(result.Formula, Clause(current))
		result.Formula.SortBySize()
		current = nil

		read++
		if read == result.Clauses {
			break
		}
	}
	return result, nil
}
