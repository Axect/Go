<h1 style="text-align:center">Answer for Problem 3</h1>
<h3 style="text-align:center">Earth's Orbit</h3>
<p style="text-align:right">Provided by <b>Tae Geun Kim</b>

## 1. Algorithm

### - Euler's Method

```LaTeX
v(t+h) ~ v(t) + h*\frac{dv(t)}{dt} = v(t)+h*a(t)
x(t+h) ~ x(t) + h*\frac{dx(t)}{dt} = x(t)+h*v(t)

h = 43200s
```

## 2. Process (Object Oriented Programming)

1. Define Const
    ```Go
    const (
        G     = 6.67259e-11       // Gravitational Constant
        m     = 5.9736e+24        // Earth mass
        M     = 1.9891e+30        // Sun mass
        AU    = 1.49597870691e+11 // Astronomy Unit
        tstep = 43200             // Time Step
        N     = 730 * 10
    )
    ```
<br>

2. Make own type - Prepare OOP
    ```Go
    type VList struct {
        R [N + 1]Vector
    }

    type Vector struct {
        x, y, z float64
    }
    ```
    <br>

    * What is OOP?
    : 자동차를 만들기 위해서는 엔진, 핸들, 바퀴, 의자, 엑셀, 브레이크 등등… 많은 부품들이 있어야 합니다. 기존 방식 절차지향적 관점으로 본다면 자동차를 만들기 위해서는 어느 한 곳을 기점으로 순서대로 만들어가야 합니다. 엔진 -> 차체 -> 핸들 -> 의자 -> 바퀴 이런 식으로 만들어가며 **이들은 서로 분리되면 안되고 순서가 틀려서도 안되며 하나가 고장 나면 전체 기능이 마비되도록 설계**되었다고 가정해봅시다. 또한 이들은 처음에 조립될 때의 부품을 다른 종류의 것으로 대체할수도 없다는 전제도 있다고 한다면! 얼마나 비효율적이고 비생산적일까요? 허나 여기에 객체지향 개념의 방식이 도입되면 상황은 달라집니다. 그러면 일단 제작에 있어서 이들은 순서적이지 않아도 됩니다. 핸들을 먼저 만들든, 바퀴를 먼저 만들든, 엔진을 먼저 만들든 상관이 없어집니다. 이들은 각각 따로 **독립적으로 개발되어 나중에 한곳에 모여 자신의 기능만 제대로 발휘**하면 되니까요. 실제 자동차를 만드는 과정과 똑같은 원리입니다. 부품들이 결합되어 움직이다 어느 하나가 고장이 나더라도 전체 부품들에 영향을 미치지는 않습니다. 고장난 부품만 고쳐주면 될 뿐, 다른 부품들은 아무 영향을 받지 않습니다. 그리고 필요하면 어느 때든 마음에 안드는 부품을 다른 것으로 교체할 수 있고 더 좋은 것으로 바꿀수도 있습니다.
    (-출처: 위시켓 블로그)
    <br>

    <img src=Fig/Object_Coffee.png>
    &nbsp; &nbsp; &nbsp; (-출처: 우아한형제들 기술 블로그)  
    <br>
    <br>

    * Advantage of OOP
        * 신뢰성 있는 소프트웨어를 쉽게 작성할 수 있다. (개발자가 만든 데이터를 사용하기에 신뢰할 수 있다.)
        * 코드를 재사용하기 쉽다.
        * 업그레이드가 쉽다. 
        * 디버깅이 쉽다.
    <br>

    * Disadvatage of OOP
    : <s>모든 야매 프로그래머들의 원수!</s>  
    데이터 클래스 개념은 상속이라는 굉장히 뛰어나지만 마찬가지로 굉장히 개 같은 특성을 지니게 해준다. 이 OOP 특성 덕분에 면밀한 자료 분석, 개발시간 단축, 좀더 정확한 코딩을 보증하지만 코드의 난이도가 급 상승한다. 한 마디로 어려워진다.  
    (-출처: 나무위키)
    <br>

3. 