File.open(ARGV[0]).each_line do |line|
    puts line.to_i.to_s(2)
end
