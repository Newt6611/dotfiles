vim.pack.add({
    { 
        src = 'https://github.com/nvim-treesitter/nvim-treesitter',
        branch = 'master'
    }
})

require'nvim-treesitter'.install { "lua", "typescript", "go", "rust" }

vim.api.nvim_create_autocmd('FileType', {
    pattern = { "lua", "typescript", "go", "rust" },
    callback = function()
        vim.treesitter.start()
    end,
})
