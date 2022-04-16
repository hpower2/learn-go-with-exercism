package tournament
import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)
type Result struct {
	wins, losses, ties int
}
type Results map[string]*Result
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	results := Results{}
	for scanner.Scan() {
		line := strings.TrimFunc(scanner.Text(), unicode.IsSpace)
		if len(line) == 0 || isComment(line) {
			continue
		}
		team, result, err := extractResult(line)
		if err != nil {
			return err
		}
		mergeResults(results, team[0], result[0])
		mergeResults(results, team[1], result[1])
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return writeResults(results, writer)
}
func isComment(line string) bool {
	r, _ := utf8.DecodeRuneInString(line)
	return r == '#'
}
func extractResult(line string) (name [2]string, res [2]Result, err error) {
	parts := strings.SplitN(line, ";", 3)
	if len(parts) != 3 {
		err = fmt.Errorf("Invalid line format: %q", line)
		return
	}
	copy(name[:], parts)
	outcome := parts[2]
	switch outcome {
	case "win":
		res[0].wins++
		res[1].losses++
	case "draw":
		res[0].ties++
		res[1].ties++
	case "loss":
		res[0].losses++
		res[1].wins++
	default:
		err = fmt.Errorf("Unknown outcome: %q", outcome)
	}
	return
}
func mergeResults(acc Results, name string, res Result) {
	current, in := acc[name]
	if !in {
		acc[name] = &res
		return
	}
	current.losses += res.losses
	current.ties += res.ties
	current.wins += res.wins
}
const tableFormat = "%-31s| %2v | %2v | %2v | %2v | %2v\n"
func writeResults(acc Results, writer io.Writer) error {
	names := namesInOrder(acc)
	w := bufio.NewWriter(writer)
	write := func(a ...interface{}) {
		fmt.Fprintf(w, tableFormat, a...)
	}
	write("Team", "MP", "W", "D", "L", "P")
	for _, n := range names {
		r := acc[n]
		mp := r.losses + r.ties + r.wins
		write(n, mp, r.wins, r.ties, r.losses, points(r))
	}
	return w.Flush()
}
func namesInOrder(acc Results) []string {
	var names []string
	for k := range acc {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		n1 := names[i]
		n2 := names[j]
		p1 := points(acc[n1])
		p2 := points(acc[n2])
		if p1 != p2 {
			return p1 > p2
		}
		return n1 < n2
	})
	return names
}
func points(r *Result) int {
	return r.wins*3 + r.ties
}