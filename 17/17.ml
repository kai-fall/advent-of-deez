open In_channel
open String
open List

let nth n l = nth l n

let lines = with_open_bin "input" input_all |> trim |> split_on_char '\n'

let a    = lines |> nth 0 |> split_on_char ' ' |> nth 2 |> int_of_string
let prog = lines |> nth 4 |> split_on_char ' ' |> nth 1 |> split_on_char ',' |> map int_of_string

let run a =
  let a = ref a in
  let b = ref 0 in
  let c = ref 0 in
  let prog = ref prog in
  let out = ref [] in
  let ptr = ref 0 in
    while !ptr < length !prog do
      let opcode = nth !ptr !prog in
      let operand = nth (!ptr + 1) !prog in
      let value =
        match operand with
          | 4 -> !a
          | 5 -> !b
          | 6 -> !c
          | _ -> operand
      in
        begin
          match opcode with
            | 0 -> a := !a lsr value
            | 1 -> b := !b lxor operand
            | 2 -> b := value mod 8
            | 3 -> if !a <> 0 then ptr := operand - 2
            | 4 -> b := !b lxor !c
            | 5 -> out := !out @ [value mod 8]
            | 6 -> b := !a lsr value
            | 7 -> c := !a lsr value
            | _ -> ()
        end;
        ptr := !ptr + 2
    done;
    !out
;;

let part1  =
  run a
  |> map string_of_int
  |> String.concat ","
  |> print_endline

let rec drop n l =
  match l with
    | []                 -> []
    | _ :: tl when n > 0 -> drop (n - 1) tl
    | _                  -> l

let part2 =
  let a = ref 0 in
    for i = (length prog - 1) downto 0 do
      a := !a lsl 3;
      while not (equal (=) (run !a) (drop i prog)) do
        a := !a + 1;
      done
    done;
    Printf.printf "%d\n" !a
;;

let () =
  part1;
  part2
