package main
import (
   "fmt"
   "math"
)

func sumVector(v1 []int, v2 []int) []int{
   result := []int{}
   
   for i := 0; i < len(v1); i++ {
      result = append(result, v1[i]+v2[i])
   }
   
   return result
}

func subVector(v1 []int, v2 []int) []int{
   result := []int{}

   for i := 0; i < len(v1); i++ {
      result = append(result, v1[i]-v2[i])
   }

   return result
}

// Interar sobre os elementos do vetor
// Elevar cada um dele ao quadrado
// Somar os elementos
// Tirar a raiz quadrada do valor obtido

// for loop
// math.Pow(2, 10)
// append

// Interar sobre o vetor v1 [x]
// Elevar cada um deles[x] ao quadrado 
// Guardar em outro vetor [x]
// Somar os valores desse novo vetor [x]

func magVector(v1 []float64) float64{
   v1sqrt := []float64{}
   resmag := 0.0
   
   for i := 0; i < len(v1); i++ {
      v1sqrt = append(v1sqrt, math.Pow(v1[i], 2))

   }

   for _, i := range v1sqrt {
         resmag = resmag + i
   }

   return  math.Sqrt(resmag) 
}

func mulVector(v1 []int, v2[]int) int {
   result := []int{}
   new_result := 0

   for i := 0; i < len(v1); i++ {
      result = append(result, v1[i]*v2[i])
      
   }

   for _, i := range result {
      new_result = new_result + i
   }

   return new_result
}

func main() {

   vec1 := []float64{1, -2, -2}
   //vec2 := []int{3, 6, 9}

   //mul := mulVector(vec1, vec2)
   //fmt.Println("Original Vector <3 -12 -18>", "Mul", mul)
   mag := magVector(vec1)
   fmt.Println(mag)

}

