from setuptools import find_packages
from setuptools import setup

setup(
    name='aburos_msgs',
    version='0.0.0',
    packages=find_packages(
        include=('aburos_msgs', 'aburos_msgs.*')),
)
