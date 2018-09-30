package main

import "time"
import "fmt"
import "math/rand"
import "sort"


const geneLength = 4


func main() {

   rand.Seed(time.Now().UnixNano())
   
   p := createPopulation(10)
   
   for i := 0; i < 1000; i++ {
      p.selection()
   }
   
   fmt.Println(p.allIndividual)
}





type gene struct {
   value int
}

func createGene() gene {
   return gene { value: rand.Intn(10) }
}

func (me *gene) mutation() {
   me.value = (me.value + rand.Intn(10 - 1)) % 10
}




type individual struct {
   allGene []gene
}

func createIndividual() individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      result.allGene[i] = createGene()
   }
   
   return result
}

func (me individual) clone() individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      result.allGene[i] = me.allGene[i]
   }
   
   return result
}

func (me *individual) fitness() int {
   fitness := 0
   for _, g := range me.allGene {
      fitness += g.value
   }
   
   return fitness
}

func (me *individual) mutation() {
   index := rand.Intn(len(me.allGene))
   me.allGene[index].mutation()
}

func crossover(allIndividual []individual) individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      index := rand.Intn(len(allIndividual))
      result.allGene[i] = allIndividual[index].allGene[i]
   }

   return result
}



type population struct {
   allIndividual []individual
}

func createPopulation(size int) population {
   result := population { allIndividual: make([]individual, size) }
   
   for i := 0; i < len(result.allIndividual); i++ {
      result.allIndividual[i] = createIndividual()
   }
   
   return result
}

func (me population) size() int { return len(me.allIndividual) }

func (me *population) selection() {
   sort.Slice(me.allIndividual, func(i, j int) bool {
      return me.allIndividual[i].fitness() < me.allIndividual[j].fitness()
   })
   
   index := rand.Intn(me.size() - 1)

   alpha := me.allIndividual[me.size() - 1]
   
   offspring := crossover([]individual { alpha, me.allIndividual[index] })
   offspring.mutation()

   me.allIndividual[index] = offspring
}





