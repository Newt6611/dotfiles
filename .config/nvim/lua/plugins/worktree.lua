return  {
    'polarmutex/git-worktree.nvim',
    version = '^2',
    dependencies = { "nvim-lua/plenary.nvim" },
    enabled = true,
    config = function ()
        require('telescope').load_extension('git_worktree')

        vim.keymap.set("n", "<leader>wt", function ()
            require("telescope").extensions.git_worktree.git_worktree()
        end)

        vim.keymap.set("n", "<leader>wc", function ()
            require('telescope').extensions.git_worktree.create_git_worktree()
        end)
    end
}
-- return {
--     'polarmutex/git-worktree.nvim',
--     version = '^2',
--     dependencies = { "nvim-lua/plenary.nvim" },
--     config = function()
--         require('telescope').load_extension('git_worktree')
--     end,
--     keys = {
--         {
--             "<leader>wt",
--             function()
--                 require("telescope").extensions.git_worktree.git_worktree()
--             end,
--         },
--         {
--             "<leader>wc",
--             function()
--                 require('telescope').extensions.git_worktree.create_git_worktree()
--             end,
--         },
--     }
-- }
