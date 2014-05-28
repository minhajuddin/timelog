#Timelog

A simple cli app which allows you to log your time

##Usage

  - Init the day by running `timelog init`
  - Log a task by running `timelog <project> <description>`, e.g.
    `timelog simplesite finished the marketing page`, where 'simplesite' will be
    parsed as the project and 'finished the marketing page' will be parsed as the
    task description
  - If you make a mistake you can open the `~/.timelog` file and edit it manually


##TODO
  - Allow initing time using `timelog init`
  - Allow logging time by running `timelog <project> <description>`
  - Allow showing the following stats
      - Hours worked today
      - Hours worked on current task
      - Other stats from github.com/minhajuddin/timelogger

  - Create a server component which allows you to analyze your time
    (copy this code from github.com/minhajuddin/timelogger)
