package main

import "time"
import "fmt"
import "math/rand"
import "sort"


const geneLength = 4


func main() {

   rand.Seed(time.Now().UnixNano())
   
   p := createPopulation(10)
   
   for i := 0; i < 100; i++ {
      p.selection()
   }
   
   p.sortFitness()
   
   for _, i := range p.allIndividual {
      fmt.Println(i)
   }
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
   fitness int
}

func createIndividual() individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      result.allGene[i] = createGene()
   }
   
   result.updateFitness();

   return result
}

func (me *individual) updateFitness() {
   me.fitness = 0
   for _, g := range me.allGene {
      me.fitness += g.value
   }
}

func crossover(allIndividual []individual, allIndex []int) individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      index := allIndex[rand.Intn(len(allIndex))]
      result.allGene[i] = allIndividual[index].allGene[i]
   }
   
   index := rand.Intn(len(result.allGene))
   result.allGene[index].mutation()
   
   result.updateFitness();
   
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

func (me *population) sortFitness() {
   sort.Slice(me.allIndividual, func(i, j int) bool {
      return me.allIndividual[i].fitness < me.allIndividual[j].fitness
   })
}

func (me *population) selection() {
   me.sortFitness()

   alphaFitness := me.allIndividual[me.size() - 1].fitness
   alphaCount := 0
   for _, i := range me.allIndividual {
      if i.fitness >= alphaFitness { alphaCount += 1 }
   }
   
   a := me.size() - rand.Intn(alphaCount) - 1
   b := rand.Intn(me.size() - alphaCount - 1)
   me.allIndividual[b] = crossover(me.allIndividual, []int { a, b })
}





