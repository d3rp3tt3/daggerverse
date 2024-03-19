import dagger
from dagger import dag, function

@function
async def lint(directory: dagger.Directory) -> str:
    return await (
        dag.container()
        .from_("python:3.10-alpine")
        .with_exec(["pip", "install", "pylint"])
        .with_mounted_directory("/src", directory)
        .with_workdir("/src")
        .with_exec(["pylint", "*.py"])
        .stdout()
    )

""" @function 
async def pyreverse(directory: dagger.Directory, format: dagger.Arg) -> str:
    return await (
        dag.container()
        .from_ ("python:3.10-alpine")
        .with_exec(["pip", "install", "pylint"])
        .with_mounted_directory("/src", directory)
        .with_workdir("/src")
        .with_exec(["pyreverse"])
        .stdout()
    )
 """

# Todo: Add parameters for the functions. Create one container reference for both functions.