### Introduction

This is a toy project to help me learn more about channels and goroutines. For this project I implemented a mini version of Hadoop which is a MapReduce framework. Instead of distributing jobs across multiple systems I'm instead distributing this across multiple goroutines.

### Future Ideas
* Figure out a better model for representing input and output to the mapper, reducer, and combiner stages.
* Combine the maximum number of goroutines (n) which will be used for both map and reduce stages. Create n buckets which will store the output of the map stage and will provide input for n reducers. Use a consistent hashing algorithm to accomodate goroutines being added / removed.