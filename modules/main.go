// 


package main
  
import (
        _"modules_learning/simpleinterest"
)

// var _ = simpleinterest.Calculate

func main() {

}

/*
The line var _ = simpleinterest.Calculate mutes the error. We should keep track of these kinds of error silencers and remove them including the imported package at the end of application development if the package is not used. Hence it is recommended to write error silencers in the package level just after the import statement.

Sometimes we need to import a package just to make sure the initialization takes place even though we do not need to use any function or variable from the package. For example, we might need to ensure that the init function of the simpleinterest package is called even though we plan not to use that package anywhere in our code. The _ blank identifier can be used in this case too as shown below.
*/