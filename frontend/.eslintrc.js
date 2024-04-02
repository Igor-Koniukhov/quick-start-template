module.exports = {
    parserOptions: {
        project: './tsconfig.json',
        tsconfigRootDir: __dirname,
        ecmaVersion: 2021,
        sourceType: 'module',
    },
    rules: {
        'react/function-component-definition': 'off',
        "import/extensions": ["error", "never", {
            "css": "always",
            "scss": "always"
        }],
        "no-underscore-dangle": "off",
        "react-hooks/exhaustive-deps": "off",
        "react/jsx-props-no-spreading": "off",
        "@typescript-eslint/type-annotation-spacing": "error",
        "react/jsx-one-expression-per-line": "off",
        "padding-line-between-statements": [
            "error",
            { "blankLine": "always", "prev": "*", "next": "block-like" },
            { "blankLine": "always", "prev": "block-like", "next": "*" }
        ]
    },
    extends: [
        'airbnb',
        'airbnb-typescript',
        'airbnb/hooks',
        'plugin:@typescript-eslint/recommended',
        'plugin:@typescript-eslint/recommended-requiring-type-checking',
    ],
    ignorePatterns: [
        '.eslintrc.js',
    ],
};
