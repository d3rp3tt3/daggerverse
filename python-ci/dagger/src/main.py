"""A simple module for testing a Node application with Dagger, written in Python.

Includes functions for building and testing a Node app. 
"""

import dagger
from dagger import dag, function, object_type


@object_type
class PythonCi:
    @function
    async def test(self, source: dagger.Directory) -> str:
        """Run unit tests"""
        return await (
            dag.node()
            .with_container(self.build_base_image(source))
            .run(["run", "test:unit", "run"])
            .stdout()
        )

    def build_base_image(self, source: dagger.Directory) -> dagger.Container:
        """Build base image"""
        return (
            dag.node()
            .with_version("12")
            .with_npm()
            .with_source(source)
            .install([])
            .container()
        )
