package cmd

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/spf13/cobra"
)

var cpuProfileFile = ""
var memoryProfileFile = ""

func Root() *cobra.Command {
	c := &cobra.Command{
		Use:   "utils",
		Short: "Has several command line utilities.",
	}
	c.AddCommand(Collatz())
	c.AddCommand(Divisors())
	c.AddCommand(Primes())
	c.AddCommand(Version())
	c.PersistentFlags().StringVarP(&cpuProfileFile, "cpu-profile", "", "", "The file to write cpu profiling data to.")
	c.PersistentFlags().StringVarP(&memoryProfileFile, "memory-profile", "", "", "The file to write memory profiling data to.")
	c.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cpuProfileFile != "" {
			cpuProfile, err := os.Create(cpuProfileFile)
			if err != nil {
				log.Fatal("failed to produce the cpu profile file: ", err)
			}
			defer cpuProfile.Close()
			if err = pprof.StartCPUProfile(cpuProfile); err != nil {
				log.Fatal("failed to start writing the cpu profile data: ", err)
			}
			defer pprof.StopCPUProfile()
		}
		if memoryProfileFile != "" {
			memoryProfile, err := os.Create(memoryProfileFile)
			if err != nil {
				log.Fatal("failed to create the memory profile file: ", err)
			}
			defer memoryProfile.Close()
			runtime.GC()
			if err = pprof.WriteHeapProfile(memoryProfile); err != nil {
				log.Fatal("failed to start writing the memory profile: ", err)
			}
		}
	}
	return c
}
