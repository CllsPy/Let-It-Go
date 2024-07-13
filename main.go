package main
import (
   "fmt"
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

// Steps:
// Criar lista vazia "new result"
// Interar sobre o vetor "result"
// A cada interação somar o valors anterior com o seguinte
// Guardar o valor dessa forma em "new_result"
// Mostrar "new_result"

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

   vec1 := []int{1, -2, -2}
   vec2 := []int{3, 6, 9}

   mul := mulVector(vec1, vec2)
   fmt.Println("Original Vector <3 -12 -18>", "Mul", mul)

}

