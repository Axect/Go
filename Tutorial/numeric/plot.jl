using Winston

Data = readcsv("Test.csv")

X = Data[:,1]
Y = Data[:,2]

x = collect(-5.:1.:5.)
y = x.^2 +2x .- 3

F = FramedPlot(
    title="Test",
    xlabel="x",
    ylabel="y"
)

P = Points(x, y, size=.1)
C = Curve(X, Y)
setattr(P, "label", "points")
setattr(C, "label", "spline")
lgnd = Legend(.5, .8, [P, C])

add(F, P, C, lgnd)
savefig(F, "Test.svg", (1000,800))

run(`inkscape -z Test.svg -e Test.png -d 300 --export-background=WHITE`)
run(`rm Test.svg`)
