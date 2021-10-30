# project-Orwell-gopher
go implementation of project orwell

## context

Project Orwell is a society simulator.

Born out of a combination of Field simulation with point objects referencing matrices of arbitrary precision.

This allows us to represent 'entities' as objects / classes, with persistent properties and locations.

It also allows us to represent the 'fabric' of a society in the form of fields, which are themselves represented as matrices.

While somewhat computationally heavy, this method allows us to generate a truly dynamic world based on randomised initial conditions.

The intention is for this to be able to function as an outline story generator.

By setting the evolution of the system as realistic, it becomes possible for authorial contrivances in the protagonists journey to have realistic consequences, and create a rich backstory which slowly reduces the need for those contrivances at all.

This version of the simulator is written in Go, as an attempt to optimise and refine the aforementioned computational intensity. 
