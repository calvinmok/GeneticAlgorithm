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

func crossover(allIndividual []individual) individual {
   result := individual { allGene: make([]gene, geneLength) }
   
   for i := 0; i < geneLength; i++ {
      index := rand.Intn(len(allIndividual))
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

   alpha := me.allIndividual[me.size() - 1]
   index := rand.Intn(me.size() / 2)
   me.allIndividual[index] = crossover([]individual { alpha, me.allIndividual[index] })
}





