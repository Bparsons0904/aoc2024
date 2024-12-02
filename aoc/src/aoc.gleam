import birl
import day1
import gleam/io

pub fn main() {
  let start_time = birl.now()
  day1.day1()
  let end_time = birl.now()

  let elapsed = birl.legible_difference(start_time, end_time)
  io.println("Execution time: " <> elapsed)
}
