// Gander provides DataFrames and Series to manipulate tabular data. It is based
// on the excellent Python Pandas package (http://pandas.pydata.org/).
// A DataFrame can be thought of as being similar to a spreadsheet, in that it holds
// rows and columns of data.
//
// Data is loaded into a DataFrame from a csv file either from
// a url, or from a file path. If all the fields of the top row of the csv contain
// non-numeric data then the top row is assumed to be column headings.
//
// Each column of the DataFrame is held as a Series object, which is made up of a
// slice of float64s, and the name of the column. Categorical (non-numeric) data
// can also be held in a Series, but no calculations can be carried out on it.
package gander
