# Bingo

I want to propose Bingo grids during our wedding. Each grid will have a number of other guests names on it.  
Then, anybody who fills one row, one column or one diagonal with guests signatures and a little story about them and us will get a small gift.  
The idea is to make people talk to each other, even if they don't know one another.

This repo allows generating bingo grids from a CSV list.

## Installation

    go install github.com/dmathieu/bingo

## Usage

You need a CSV list of names as input. The program will generate a PDF document as output.

    bingo -in my_list.csv -out my_document.pdf -size 7

The size argument is the number of entries there will be on each row, column or diagonal.  
Every grid is forced to be a square, or some ways would be easier than others.

The program will put every entry in random order the same number of time.  
In order to balance that, there can be a lot of pages in your grid, up to as many entries as you have.
The same entry cannot be put on the page twice.

## Known issues

### "Too many retries"

There is a possibility that the very last entry has nowhere to be put to (if it is already on the page holding the last available element).  
When this happens, a "too many retries" error will happen.

There are a few ways we could fix this. But I'm planning a wedding goddamit, I can't do everything!  
So, just rerun the thing to get something valid.

## License

Bingo is released under the [MIT License](http://www.opensource.org/licenses/MIT).
