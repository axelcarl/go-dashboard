import * as z from "zod";

export const PaymentSchema = z
  .object({
    id: z.number(),
    sender: z.string(),
    recipient: z.string(),
    amount: z.number(),
    createdAt: z.string(),
    updatedAt: z.string(),
  })
  .transform((payment) => ({
    ...payment,
    createdAt: new Date(payment.createdAt),
    updatedAt: new Date(payment.updatedAt),
  }));

export type Payment = z.infer<typeof PaymentSchema>;

export const PaymentsSchema = z.array(PaymentSchema);
export type Payments = z.infer<typeof PaymentsSchema>;
