sym = [ ["-**--", "*--*-", "*--*-", "*--*-", "-**--", "-----"],
        ["--*--", "-**--", "--*--", "--*--", "-***-", "-----"],
        ["***--", "---*-", "-**--", "*----", "****-", "-----"],
        ["***--", "---*-", "-**--", "---*-", "***--", "-----"],
        ["-*---", "*--*-", "****-", "---*-", "---*-", "-----"],
        ["****-", "*----", "***--", "---*-", "***--", "-----"],
        ["-**--", "*----", "***--", "*--*-", "-**--", "-----"],
        ["****-", "---*-", "--*--", "-*---", "-*---", "-----"],
        ["-**--", "*--*-", "-**--", "*--*-", "-**--", "-----"],
        ["-**--", "*--*-", "-***-", "---*-", "-**--", "-----"] ]
File.open(ARGV[0]).each_line do |line|
    digits = line.gsub(/\D/, '').chars.map(&:to_i)
    (0...sym[0].length).each {|r| puts digits.map {|d| sym[d][r]}.join('')}
end
