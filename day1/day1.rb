p File.read('./input.txt').split("\n").map{|x| x.to_i}.permutation(2).to_a.select{|x,y| x + y == 2020}.first.reduce(:*)
