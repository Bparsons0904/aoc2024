import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub fn day1() {
  read_file("../inputs/input1a.txt")
}

pub fn read_file(path: String) {
  case simplifile.read(path) {
    Ok(content) -> {
      content
      |> string.split("\n")
      |> process_lines
    }
    Error(error) -> {
      io.println(simplifile.describe_error(error))
    }
  }
}

pub fn process_lines(lines: List(String)) {
  let line_pairs =
    list.map(lines, fn(line) {
      case string.split(line, "   ") {
        [first, second] -> {
          let a = int.parse(string.trim(first)) |> result.unwrap(0)
          let b = int.parse(string.trim(second)) |> result.unwrap(0)
          #(a, b)
        }
        _ -> #(0, 0)
      }
    })

  let firsts =
    list.map(line_pairs, fn(pair) { pair.0 })
    |> list.sort(int.compare)

  let seconds =
    list.map(line_pairs, fn(pair) { pair.1 })
    |> list.sort(int.compare)

  let difference = count_difference(firsts, seconds, 0)
  io.println(int.to_string(difference))

  let total = count_total(firsts, seconds, 0)
  io.println(int.to_string(total))
}

fn count_total(first_list: List(Int), second_list: List(Int), total: Int) -> Int {
  case first_list {
    [] -> total
    [current, ..rest] -> {
      let occurrences =
        list.filter(second_list, fn(num) { num == current })
        |> list.length()

      case occurrences {
        0 -> count_total(rest, second_list, total)
        _ -> count_total(rest, second_list, total + current * occurrences)
      }
    }
  }
}

fn count_difference(
  first_list: List(Int),
  second_list: List(Int),
  total: Int,
) -> Int {
  case first_list, second_list {
    [], [] -> total
    [a, ..rest_a], [b, ..rest_b] -> {
      let diff = int.absolute_value(a - b)
      count_difference(rest_a, rest_b, total + diff)
    }
    _, _ -> total
  }
}
