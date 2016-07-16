'use strict';

const gulp = require('gulp'),
    config = require('../config');
    
    
gulp.task('build', ['{{.npm.language}}'])