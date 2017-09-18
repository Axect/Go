using Winston

Data = readcsv("sin.csv")

x = Data[:,1]
y = Data[:,2]

C = Curve(x, y, color="black")
setattr(C, "label", "Sin")
lgnd = Legend(.9, .9, [C])

pl = FramedPlot(
    title="Sin Graph",
    xlabel="x",
    ylabel="y"
)

add(pl, C, lgnd)
savefig(pl, "sin.svg", (800, 800))
run(`inkscape -z sin.svg -e sin.png -d 300 --export-background=WHITE`)
run(`rm sin.svg`)