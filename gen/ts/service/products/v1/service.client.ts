// @generated by protobuf-ts 2.2.2 with parameter long_type_string
// @generated from protobuf file "service/products/v1/service.proto" (package "service.products.v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { ProductService } from "./service";
import type { AllProductsResponse } from "./service";
import type { AllProductsRequest } from "./service";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ProductServiceCreateProductResponse } from "./service";
import type { CreateProductRequest } from "./service";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service service.products.v1.ProductService
 */
export interface IProductServiceClient {
    /**
     * @generated from protobuf rpc: CreateProduct(service.products.v1.CreateProductRequest) returns (service.products.v1.ProductServiceCreateProductResponse);
     */
    createProduct(input: CreateProductRequest, options?: RpcOptions): UnaryCall<CreateProductRequest, ProductServiceCreateProductResponse>;
    /**
     * @generated from protobuf rpc: AllProducts(service.products.v1.AllProductsRequest) returns (service.products.v1.AllProductsResponse);
     */
    allProducts(input: AllProductsRequest, options?: RpcOptions): UnaryCall<AllProductsRequest, AllProductsResponse>;
}
/**
 * @generated from protobuf service service.products.v1.ProductService
 */
export class ProductServiceClient implements IProductServiceClient, ServiceInfo {
    typeName = ProductService.typeName;
    methods = ProductService.methods;
    options = ProductService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: CreateProduct(service.products.v1.CreateProductRequest) returns (service.products.v1.ProductServiceCreateProductResponse);
     */
    createProduct(input: CreateProductRequest, options?: RpcOptions): UnaryCall<CreateProductRequest, ProductServiceCreateProductResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateProductRequest, ProductServiceCreateProductResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: AllProducts(service.products.v1.AllProductsRequest) returns (service.products.v1.AllProductsResponse);
     */
    allProducts(input: AllProductsRequest, options?: RpcOptions): UnaryCall<AllProductsRequest, AllProductsResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<AllProductsRequest, AllProductsResponse>("unary", this._transport, method, opt, input);
    }
}
