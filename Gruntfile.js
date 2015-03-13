'use strict';

module.exports = function(grunt) {

    
    grunt.initConfig({
        watch:{                      
            project: {
                files:['application/**/*.go','core/**/*.go'],
                tasks:['shell:stop','shell:start']
            }                       
        },
        shell: {            
                       
            start: {
                command: [
                    'go run index.go &',
                    'grunt watch:project'                 
                    ].join('')
            },
            stop: {
                command: 'pkill -9 index &'
            }            
        }

    });

    grunt.loadNpmTasks('grunt-contrib-watch');
    grunt.loadNpmTasks('grunt-shell');

    grunt.registerTask('start', ['shell:start']);
    grunt.registerTask('stop', ['shell:stop']);


};