'use strict';

const gulp = require('gulp'),
    config = require('../config'),
    tsc = require('gulp-typescript'),
    merge = require('merge2');


var project = tsc.createProject('./tsconfig.json', {

});

gulp.task('typescript', ['addfiles'], function () {

    let result = project.src()
        .pipe(tsc(project));

    return merge([
        result.js.pipe(gulp.dest(config.outputDir)),
        result.dts.pipe(gulp.dest(config.outputDir))
    ]);
});

var fs = require('fs');
var readdir = require('recursive-readdir');

gulp.task('addfiles', function () {
    var tsconfigDir = process.cwd() + '/tsconfig.json';

    var tsconfig = require(tsconfigDir);

    readdir(process.cwd() + '/src', function (e, files) {
        tsconfig.files = files.filter(function (file) {
            var len = file.length;
            //if (file.indexOf('tools') > -1) return false;
            return (file.substr(len - 3) === '.ts' || file.substr(len - 3) === '.js')
                && file.substr(len - 5) !== ".d.ts";
        }).map(function (file) {
            return file.replace(process.cwd() + '/', '');
        });

        tsconfig.files.unshift('typings/index.d.ts')

        fs.writeFile(tsconfigDir, JSON.stringify(tsconfig, null, 2), function () {
            console.log('%s files added', tsconfig.files.length);
            done();
        });
    });
});