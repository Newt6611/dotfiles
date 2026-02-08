
local o = vim.opt
local g = vim.g

o.termguicolors = true

o.number = true
o.relativenumber = true

o.hlsearch = false
o.swapfile = false

o.clipboard = 'unnamedplus'
o.winborder = 'rounded'
o.signcolumn = 'yes'

-- Number of spaces that a <Tab> in the file counts for
o.tabstop = 4
-- Number of spaces to use for each step of (auto)indent
o.shiftwidth = 4
-- Convert tabs to spaces automatically
o.expandtab = true
-- Number of spaces that a <Tab> counts for while performing editing operations
o.softtabstop = 4

g.mapleader = " "
