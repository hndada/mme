# MME
Minimal Markup Engine

# Moo
A compact library for constructing web application with Go
* Build web components by Go struct `Block`.
    * Styles are structured 
* Handle components behaviors by implementing Block's methods.
    * When to re-render
    * When to change styles
* It outputs to proper html files.
    * User doesn't have to know how html or javascript work.

# Proposal
* Implement as div element. Output as semantic tags.
    * Compare the difference between initial properties and attributes 
    then apply to semamntic tags when output 

# Dropped properties
The common reason that properties go dropped is that 
it requires several linking to see the actual value.

* all
* appearance
* aspect-ratio: controlled by script
* backface-visibility: too advanced at this moment
* break: for printing?
* caption-side: can be handled in more general ways

# Not-visited yet
* align
