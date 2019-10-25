set terminal png size 1024,780

set output "performance.png"

set title "Performance test"

set grid y

set xlabel "Requests"

set ylabel "Time (ms)"

plot "nodejs.tsv" using 9 smooth sbezier with lines title "NodeJS", \
     "golang.tsv" using 9 smooth sbezier with lines title "GoLang", \
     "netcore3.tsv" using 9 smooth sbezier with lines title "Net Core 3"

exit