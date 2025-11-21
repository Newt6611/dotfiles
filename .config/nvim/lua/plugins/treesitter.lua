return {
    "nvim-treesitter/nvim-treesitter",
    opts = {
        ensure_installed = {
            "html",
            "lua",
            "luadoc",
            "luap",
            "toml",
            "typescript",
            "vim",
            "yaml",
            "go",
            "rust",
        },
        indent = {
            enable = true,
        }
    },
}
