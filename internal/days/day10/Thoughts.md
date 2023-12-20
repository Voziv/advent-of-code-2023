# Data structure

My first instinct was to read this in as a 2d array `var pipeMaze [][]string` but the main issue is accessing the coords
is backwards: `pipeMaze[y][x]`. I personally find this frustrating and want to find a way to access this via x,y.

Drew on the kdev discord did remind me that typically in gamedev it's common to store the data in a linear array rather
than a 2d array. Given we're not concerned with memory size here this should work great!

# Part two

At first, I tried a flood fill. This could work if I expanded cells to be larger so that we can iterate between
everything. Instead, I've opted to scan line by line checking boundaries to figure out what's inside and outside the
shape.

Another option is using the shoelace and pickers theorem. I might come back and re-implement someday.... so never ;)

I did get a nicely [rendered map](RenderedMap.md) out of it.