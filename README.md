## Introduction
**Ray Tracing** is an algorithm used to create realistic graphics.

## Sub-problem No. 1
Any Ray Tracing algorithm requires the representation of objects in a 3D space.

A vector is a line segment that extends from point A to point B in a 2D or 3D plane.

We need to be able to operate on vectors.

Normalizing a vector involves scaling its magnitude/size to 1 while retaining its direction. This can be useful for other operations such as distance calculation or clustering.

The formula for the cross product is given by:
![image](image_2.png)

To implement the calculation, the following steps were taken:

- Create an empty list called `new_result`
- Iterate over the vector `result`
- In each iteration, add the previous value to the next one
- Store this result in `new_result`
- Display `new_result`

To calculate the magnitude of a vector:

### Steps
- Iterate over the vector elements
- Square each element
- Sum the squared values
- Take the square root of the resulting value

### Resources
- `for` loop
- `math.Pow(2, 10)`
- `append`

### Plan
- Iterate over the vector `v1[x]`
- Square each `v1[x]`
- Store the squared values in another vector `[x]`
- Sum the values in this new vector `[x]`

## Useful Links
- [Tour for Go](https://go.dev/tour/moretypes/6) ![Go Tour](https://img.shields.io/badge/Go_Tour-%2300ADD8.svg?style=flat-square&logo=go)
- [Vector Representation in Go](https://www.netguru.com/blog/vector-operations-in-go) ![Netguru Blog](https://img.shields.io/badge/Netguru_Blog-%2300B9F2.svg?style=flat-square&logo=netguru)
- [Vector Normalization](https://www.khanacademy.org/computing/computer-programming/programming-natural-simulations/programming-vectors/a/vector-magnitude-normalization) ![Khan Academy](https://img.shields.io/badge/Khan_Academy-%23F2B000.svg?style=flat-square&logo=khan-academy)
- [OOP in Golang](https://www.tutorialspoint.com/golang-program-to-create-a-class-and-object#:~:text=Structs%20%E2%88%92%20Go%20language%20does%20not,same%20or%20different%20data%20type.) ![TutorialsPoint](https://img.shields.io/badge/TutorialsPoint-%236D0F8C.svg?style=flat-square&logo=tutorialspoint)
- [Machine Learning with Go: Matrices and Vectors | packtpub.com](https://www.youtube.com/watch?v=rzYzsdKImEs) ![PacktPub](https://img.shields.io/badge/PacktPub-%23B4006F.svg?style=flat-square&logo=youtube)
- [Sum Vector JS Solution](https://spellbox.app/spells/javascript/two-functions-that-make-sum-two-vectors-and-make-their-dot-product-in-javascript) ![Spellbox](https://img.shields.io/badge/Spellbox-%23F8E71C.svg?style=flat-square&logo=spellbox)
- [Golang - Append Function](https://dev.to/andyhaskell/a-closer-look-at-go-s-slice-append-function-3bhb) ![Dev.to](https://img.shields.io/badge/Dev.to-%23F7DF1E.svg?style=flat-square&logo=dev)
- [Passing & Returning Array To & From a Function | Go Tutorial](https://www.youtube.com/watch?v=_mVLhNgQ7_8) ![YouTube](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=flat-square&logo=youtube)
- [The Vector Product](https://www.mathcentre.ac.uk/resources/uploaded/mc-ty-vectorprod-2009-1.pdf) ![Math Centre](https://img.shields.io/badge/Math_Centre-%2318C5C5.svg?style=flat-square&logo=math)
