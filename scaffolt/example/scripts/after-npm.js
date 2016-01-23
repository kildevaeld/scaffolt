
var _ = require('underscore');

function initTypescript (context, pkgjson) {
    _.extend(pkgjson.dependencies, {
        tslint: "*"
    });
    
    if (pkgjson.build === 'gulp') {
        _.extend(pkgjson.devDependencies, {
            'gulp-typescript': '*',
            'merge2': '*'    
        });
        //context.Move('src/gulp/tasks/typescript.js', 'gulp/tasks/typescript.js', true);
    
    }
    //context.Move('src/tsconfig.json', 'tsconfig.json', true);
    
}

function initGulp (context, pkgjson) {
    _.extend(pkgjson.devDependencies, {
       gulp: "*",
       requiredir: "*" 
    });
   
   context.Move("src/gulp/gulpfile.js", "gulpfile.js", true);
    //context.CreateFile('gulpfile.js', '"use strict";\nrequire("requiredir")(gulp/tasks", {recursive: true});')
    
    //context.Move('src/gulp/tasks/build.js', 'gulp/tasks/build.js', true);
    //context.Move('src/gulp/tasks/default.js', 'gulp/tasks/default.js', true);
    //context.Move('src/gulp/config.js', 'gulp/config.js', true);
}

function initTest(context, pkgjson) {
    switch (pkgjson.test) {
        case "mocha":
            _.extend(pkgjson.devDependencies, {
                mocha: "*",
                should: "*"
            });
            pkgjson.scripts.test = "node node_modules/.bin/mocha";
    }
}


module.exports = function (context) {
    
    var npm = context.Get("npm");
    var pkg = _.extend({},npm, {
        dependencies:{},
        devDependencies:{},
        scripts: {},
        main: "index.js"
    });
    
    switch (npm.language) {
        case "typescript":
            initTypescript(context, pkg)
    }
    
    switch (npm.type) {
        case "nodejs":
        case "browser":
    }
    
    switch (npm.build) {
        case "gulp":
            initGulp(context, pkg);
            break;
        case "grunt":
    }
    
    if (npm.test != "none") {
        initTest(context, pkg)
    }
    
    pkg = _.omit(pkg, ['build', 'language', 'type']);
    
    var json = JSON.stringify(pkg, null, "  ");
    
    context.CreateFile('package.json', json);
    context.CreateFile('.gitignore', 'node_modules\n.DS_Store\n*.log')
    
    context.Exec("npm", "install", '--save', '--save-dev');
}