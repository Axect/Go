<h1 style='text-align:center'> Problem 2 with Go</h1>
<p style="text-align:right"> Provided by <b> Tae Geun Kim </b></p>

## 1. Roll Dice

1. Make Roll function :``` Input-int, output-[]int64 (The number of rolling, Result)```
    ```Go
    func Roll(n int) []int64 {
	    Result := make([]int64, n, n)
	    rand.Seed(int64(time.Now().Nanosecond()))
	    for i := range Result {
		    results := rand.Intn(6) + 1
		    Result[i] = int64(results)
	    }
	    return Result
    }
    ```
    * rand.Seed(n int64) : There is no true random. Thus should use Seed
    * rand.Intn(n int) : From 0 to n-1 random number
    * Slice can change value (Result[i] = ~)
<br>

2. Make Quincunx function :``` Input-[]int64, output-float64 (Result, probability)```
    ```Go
    func Quincunx(array []int64) float64 {
	    length := float64(len(array))
	    count := 0
	    for _, elem := range array {
		    if elem == 5 {
			    count++
		    }
	    }
	    return float64(count) / length
    }
    ```
    * Go has Auto enumerate : ```for _, elem := range array``` - index, element

## 2. Sort List

1. Make some types
    ```Go
    // PhysicsTest includes name & score
    type PhysicsTest struct {
	    Name  string
	    Score int
    }

    // String() is just string method
    func (p PhysicsTest) String() string {
	    return fmt.Sprintf("%s: %d", p.Name, p.Score)
    }

    // PhysicsList : Let's type to list
    type PhysicsList []PhysicsTest

    // NewContainer contains rank
    type NewContainer struct {
	    Inform PhysicsTest
	    Rank   int
    }
    ```
    * PhysicsTest : ```A := PhysicsTest{"Schrodinger", 66}``` $\Rightarrow$ ```Print(A.Name) -> "Schrodinger"```
    * String() : method for PhysicsTest ```Print(A) ->  "Schrodinger": 66 ```
    * PhysicsList : Array of PhysicsTest
    * NewContainer : PhysicsTest + Rank
    * Example:
        ```Go
        func main() {
            A := PhysicsTest{"Schrodinger", 66}
            B := NewContainer{A, 6}
            fmt.Println((B.Inform, B.Rank))
        }
        ```
        **Q.** What's the output?
<br>

2. Method for PhysicsList

    ```Go
    func (p PhysicsList) Len() int           { return len(p) }
    func (p PhysicsList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    func (p PhysicsList) Less(i, j int) bool { return p[i].Score < p[j].Score }

    // Mean is mean
    func (p PhysicsList) Mean() float64 {
        s := 0
        l := p.Len()
        for _, elem := range p {
            s += elem.Score
        }
        return float64(s) / float64(l)
    }

    // Std is std
    func (p PhysicsList) Std() float64 {
        s := float64(0)
        l := p.Len()
        m := p.Mean()
        for _, elem := range p {
            s += math.Pow(float64(elem.Score)-m, 2)
        }
        return math.Sqrt(s / (float64(l) - 1)) // -1 means sample
    }
    ```
    * Len, Swap, Less : Necessary method for sorting (ex: ```A.Len() : length of A ```)
    * Mean, Std : As you know. But let's see more specifics.
        ```Go
        for _, elem := range p {
            s += elem.Score
        }
        ```
        * ```p : PhysicsList ([]PhysicsTest)```
        * ```elem : PhysicsTest (which has Name & Score)```
<br>

3. Method for NewContainer

    ```Go
    func (Elem NewContainer) String() string {
        return fmt.Sprintf("%s, %d, %d\n", Elem.Inform.Name, Elem.Inform.Score, Elem.Rank)
    }

    // AssignRank - rank
    func (Elem *NewContainer) AssignRank(n int) { Elem.Rank = n }

    // AssignInform - inform
    func (Elem *NewContainer) AssignInform(value PhysicsTest) { Elem.Inform = value }

    // Rank ranks list
    func Rank(SortedArray PhysicsList) []NewContainer {
        length := SortedArray.Len()
        Temp := make([]NewContainer, length, length)
        for i := range SortedArray {
            if i > 0 && SortedArray[i].Score == SortedArray[i-1].Score {
                Temp[i].AssignRank(int(Temp[i-1].Rank))
            } else {
                Temp[i].AssignRank(int(i + 1))
            }
            Temp[i].AssignInform(SortedArray[i])
        }
        return Temp
    }
    ```
    * ```(Elem NewContainer) : Normal Receiver - Does not change value```
    * ```(Elem *NewContainer) : Pointer Receiver - Change Value```
        ```Go
        func main() {
            A := PhysicsTest{"Schrodinger", 66}
            var B NewContainer
            B.AssingInform(A)
            B.AssignRank(6)
            fmt.Println(B.Inform.Name, B.Rank)
        }
        ```
    * How to handle Joint Rank? : 
        * ```SortedArray : PhysicsList ([]PhysicsTest)```
        * ```SortedArray[i] : PhysicsTest (has Name, Score) ```
        * ```SortedArray[i] .Score == SortedArray[i-1].Score : Same Score (Joint)```
        * ```Temp : []NewContainer -> Temp[i] : NewContainer```
        * ```Temp[i].AssignRank(int(Temp[i-1].Rank)) : Same rank assign``` 
<br>

4. Ready to write CSV
    ```Go
    // ConvertString converts string list
    func ConvertString(C []NewContainer) [][]string {
        length := len(C)
        Temp := make([][]string, length, length)
        for i, elem := range C {
            Temp[i] = []string{fmt.Sprint(elem.Inform.Name), fmt.Sprint(elem.Inform.Score), fmt.Sprint(elem.Rank)}
        }
        return Temp
    }

    // -----------------------------
    // Package/csv/csv.go
    // -----------------------------

    // Write writes csv file
    func Write(List [][]string, name string) {
        Title := name
        file, err := os.Create(Title)
        checkError("Cannot create a file", err)
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        for _, value := range List {
            err := writer.Write(value)
            checkError("Cannot write to file", err)
        }
        fmt.Println("Complete to Write")
    }

    func checkError(message string, err error) {
        if err != nil {
            log.Fatal(message, err)
        }
    }
    ```
    * I made csv package (contain Read & Write & CheckError)
        * Write : ```Input: [][]string, string -> Output: Nothing but .csv file```
        * CheckError : ```Input: string, error -> Output: Nothing but exit```
    * Thus make ```ConvertString``` function
        * ```Input: []NewContainer -> Output: [][]string```
        * ```fmt.Sprint: convert any value to string```
<br>

5. Main Function
    ```Go
    // DoSort do sort
    func DoSort() {
        fmt.Println("---------------------------------------")
        fmt.Println("Sort List")
        fmt.Println("---------------------------------------")
        A := PhysicsList{
            {"Schrodinger", 64},
            {"Einstein", 90},
            {"Feynman", 98},
            {"Neumann", 100},
            {"Dirac", 92},
            {"Bohr", 90},
            {"Fermi", 91},
            {"Heisenberg", 72},
            {"Pauli", 88},
            {"Newton", 93},
            {"Leibniz", 93},
            {"Planck", 52},
        }

        // Sorting
        sort.Sort(sort.Reverse(A))

        // Statistic
        fmt.Printf("Average: %f\n", A.Mean())
        fmt.Printf("Std: %f\n ", A.Std())

        // Add Rank
        Result := Rank(A)

        // Write CSV
        Data := ConvertString(Result)
        csv.Write(Data, "../Data/physics.csv")

        // Read CSV
        List := csv.Read("../Data/physics.csv")
        Physicsists := make([]string, len(List), len(List))
        Scores, Ranks := make([]int64, len(List), len(List)), make([]int64, len(List), len(List))
        for i, group := range List {
            Physicsists[i] = group[0]
            Scores[i], _ = strconv.ParseInt(group[1], 0, 64)
            Ranks[i], _ = strconv.ParseInt(group[2], 0, 64)
        }
        fmt.Println(Ranks)
    }
    ```
    * Read : ```Input: csv -> Output: [][]string ```
<br>

6. Plotting in Python
    ```Python
    # ==============================================================================
    # Python Import Session
    # ==============================================================================

    import matplotlib.pyplot as plt
    import seaborn

    # ==============================================================================
    # Receive Sequence
    # ==============================================================================

    temp = open("Data/physics.csv")
    List = [line.split(',') for line in temp]
    Name, Score, Rank = [], [], []
    for group in List:
        Name.append(group[0])
        Score.append(int(group[1]))
        Rank.append(int(group[2]))

    # ==============================================================================
    # Plotting
    # ==============================================================================

    plt.figure(figsize=(10,10), dpi=600)
    plt.plot(Rank, Score, label='Score')
    plt.title("Distribution of Score")
    plt.xlabel('Rank')
    plt.ylabel('Score')
    plt.legend()
    plt.savefig('Fig/Distrib.png')
    ```
    * Read is so simple : ```open & line.split by list comprehension```
    <img src="Fig/Distrib.png"></img>