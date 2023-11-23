const path = require('path');

const {WebpackManifestPlugin} = require('webpack-manifest-plugin');
const TerserPlugin = require("terser-webpack-plugin");
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyPlugin = require("copy-webpack-plugin");

const host = process.env.HOST || 'localhost';
const devServerPort = 3808;

const production = process.env.NODE_ENV === 'production';

const config = {
    mode: production ? 'production' : 'development', resolve: {
        extensions: ['.js', '.css'],
    }, entry: {
        "public": ['./static/js/index.js', './static/css/base.css',],
        "application": ['./static/js/index.js', './static/css/base.css',],
    }, devServer: {
        static: {directory: path.resolve(__dirname, 'dist')},
        port: devServerPort,
        open: false,
        hot: true,
        compress: true,
        historyApiFallback: true,
    }, output: {
        path: path.resolve(__dirname, 'dist'),
        filename: production ? 'js/[name]-[chunkhash].js' : 'js/[name].js',
        chunkFilename: production ? 'js/[name]-[chunkhash].js' : 'js/[name].js',
        publicPath: '/dist',
        clean: true,
    }, module: {
        rules: [{
            test: /\.js$/, exclude: /(node_modules)/, use: {
                loader: 'babel-loader', options: {
                    presets: ['@babel/preset-env'],
                },
            },
        }, {
            test: /\.css$/i, use: [{loader: MiniCssExtractPlugin.loader}, {
                loader: 'css-loader', options: {sourceMap: true}
            }, 'postcss-loader',],
        },],
    }, plugins: [new MiniCssExtractPlugin({
        filename: production ? "css/[name]-[chunkhash].css" : "css/[name].css",
        chunkFilename: production ? "css/[name]-[id].css" : "css/[name].css",
    }), new WebpackManifestPlugin({
        writeToFileEmit: true, publicPath: production ? "/dist/" : 'http://' + host + ':' + devServerPort + '/dist/',
    }), new CopyPlugin({
        patterns: [{from: "./static/robots.txt"}],
    }),], optimization: {
        minimize: production, minimizer: [new TerserPlugin(), new CssMinimizerPlugin(),],
    }
};

if (!production) {
    config.devServer = {
        port: devServerPort, headers: {'Access-Control-Allow-Origin': '*'},
    };

    config.output.publicPath = 'http://' + host + ':' + devServerPort + '/dist/';
    // Source maps
    config.devtool = 'source-map';
}

module.exports = config