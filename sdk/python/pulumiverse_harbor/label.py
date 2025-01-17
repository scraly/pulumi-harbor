# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LabelArgs', 'Label']

@pulumi.input_type
class LabelArgs:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Label resource.
        """
        LabelArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            color=color,
            description=description,
            name=name,
            project_id=project_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             color: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if color is not None:
            _setter("color", color)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _LabelState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Label resources.
        """
        _LabelState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            color=color,
            description=description,
            name=name,
            project_id=project_id,
            scope=scope,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             color: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             scope: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']

        if color is not None:
            _setter("color", color)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if scope is not None:
            _setter("scope", scope)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


class Label(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        * Create a global label within harbor

        * Creates a label for project

        ## Import

        Harbor label can be imported using the `label id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/label:Label main /labels/1 <break>```<break><break>`

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LabelArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        * Create a global label within harbor

        * Creates a label for project

        ## Import

        Harbor label can be imported using the `label id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/label:Label main /labels/1 <break>```<break><break>`

        :param str resource_name: The name of the resource.
        :param LabelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LabelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            LabelArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LabelArgs.__new__(LabelArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["scope"] = None
        super(Label, __self__).__init__(
            'harbor:index/label:Label',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'Label':
        """
        Get an existing Label resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LabelState.__new__(_LabelState)

        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["scope"] = scope
        return Label(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        return pulumi.get(self, "scope")

