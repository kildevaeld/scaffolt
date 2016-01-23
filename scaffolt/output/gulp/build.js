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

gulp.task('addfiles', function () {
    
});